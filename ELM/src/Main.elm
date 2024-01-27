module Main exposing (..)
import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http exposing (Error(..))
import List exposing (head, length, drop, take)
import Json.Decode as JSONDecode exposing (Decoder)
import Random



-------- MAIN ---------------------

main : Program () Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }

-- Types
type State 
  = Loading
  | Error String
  | Guessing


type alias Definition =
    { partOfSpeech : String
    , defs : List String
    }
  
-- MODEL

type alias Model  = 
  { dicoList : List String
  , definitions : List Definition
  , guess : String
  , guessWord : String
  , state : State
  , title : String
  , question : String
  }
 

init : () -> (Model, Cmd Msg)
init _ = 
    ( { dicoList = []
    , definitions = []
    , guessWord  = ""
    , guess = ""
    , state = Loading
    , title = "Guess it!"
    , question = "Type in to guess"
    }
    , Http.get { url = "dico.txt" 
                , expect = Http.expectString GotText}

    )


-- UPDATE

type Msg
  = GotText (Result Http.Error String)
  | GotWord String
  | GotDefJson (Result Http.Error (List Definition))
  | Guessed String
  | ShowAnswer
  | HideAnswer
  



update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    -- Quand on a obtenu le contenu de dico.txt
    GotText result ->
      case result of
        Ok fullText -> 
          let 
            dicoList = String.words fullText --recupere une liste de mot
          in 
          case dicoList of 
            [] ->  ({model | state = Error "Empty"}
                   , Cmd.none) --on change le statut du modèle en erreur

            x::xs -> ({model | dicoList = dicoList}
                   , getRandomElement x xs ) --Si la liste n'est pas vide, on cherche un mot au hasard

        Err _ -> ({model | state = Error "I can't load your words"},
                     Cmd.none)

    -- Quand on a le mot, on met à jour le modele et on cherche sa définition
    GotWord word -> ({model | guessWord = word}, getDefinition word)


    --Quand on a extrait la bonne définition du json
    GotDefJson  result ->
      case result of 
        Ok definitions -> ({model | definitions = definitions, state = Guessing}
                          , Cmd.none)
        Err _ -> ({model | state = Error "I can't load definitions"}
                  , Cmd.none)
    -- Quand le joueur écrit dans la zone de texte
    Guessed guess ->
      if guess == model.guessWord then
             ({model | guess = guess
             , question = ("Got it! It is indeed "++ model.guessWord)}
              , Cmd.none)
      else
             ({model | guess = guess, question = "Type in to guess"}
              , Cmd.none)
    -- Quand le joueur demande de montrer ou de cacher la réponse.
    ShowAnswer ->  ({model | title = model.guessWord }
              , Cmd.none)
    HideAnswer ->  ({model | title = "Guess it!" }
              , Cmd.none)
        
-- Fonction permettant d'extraire un mot aléatoire d'une liste de mot et d'un mot
getRandomElement : String -> List String -> Cmd Msg
getRandomElement x xs =
  Random.generate GotWord (Random.uniform x xs)

-- Fonction permmettant de récupérer la définition sous format json d'un mot.
getDefinition : String -> Cmd Msg
getDefinition word =
  let 
    urlApi = "https://api.dictionaryapi.dev/api/v2/entries/en/"
  in 
  Http.get { url = urlApi ++ word , expect = Http.expectJson GotDefJson jsonDecoder
      }

-- Extrait la definition : string, de la clef definition
sentenceDecoder : Decoder String
sentenceDecoder =
  JSONDecode.field "definition" JSONDecode.string

-- map2 : combiner les résultats des deux décodeurs :
-- le premier extrait la valeur de "partOfSpeech"
-- le second extrait une liste de definition contenue dans "definitions"

definitionDecoder : Decoder Definition
definitionDecoder =
    JSONDecode.map2 Definition
        (JSONDecode.field "partOfSpeech" JSONDecode.string)
        (JSONDecode.field "definitions" (JSONDecode.list sentenceDecoder))

-- Accède au chemin ["0", "meanings"]  (.at)
jsonDecoder : Decoder (List Definition)
jsonDecoder =
  JSONDecode.at ["0", "meanings"](JSONDecode.list definitionDecoder)


      
-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW
view : Model -> Html Msg
view model =
    let
        viewDef line =
            li [ Html.Attributes.style "font-style" "normal"] [ text line ]

        viewMeaning meaning =
            ul []
                [ li [ Html.Attributes.style "font-style" "italic"]
                    [ text meaning.partOfSpeech
                    , ol [] (List.map viewDef meaning.defs) 
                    ] --ol liste avec numéro List.map applique viewDef sur toutes les lignes
                ]
        guessingView =
            div []
                [ h1 [ Html.Attributes.style "font-family" "Didot" ] [ text model.title ]
                , h4 [Html.Attributes.style "font-family" "Didot" ] [ text "Meanings:" ]
                , ul [Html.Attributes.style "font-family" "Didot"] (List.map viewMeaning model.definitions)
                , h3 [Html.Attributes.style "font-family" "Didot"] [ text model.question ]
                , input [ onInput Guessed ] []
                , div []
                  [button [ onClick ShowAnswer ] [ text " Show answer" ]
                , button  [ onClick HideAnswer] [ text " Hide answer" ]
                  ]
                ]
    in
    case model.state of
        Loading ->
            div [] []

        Error errorMsg ->
            h3 [] [ text ("Error: " ++ errorMsg) ]

        Guessing ->
            guessingView

    
