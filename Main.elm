import Browser
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, map4, field, int, string)



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL

type alias Model =
  { content : String
  }

init : Model
init =
  { content = "" }


type alias Word =
  { content : String
  }

init : () -> (Model, Cmd Msg)
init _ =
  (Loading, randomWord)



-- UPDATE

type Input
  = Change String


update : Input -> Model -> Model
update Input model =
  case Input of
    Change newContent ->
      { model | content = newContent }

type Msg
  = GotText (Result Http.Error String)

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          (Success fullText, Cmd.none)

        Err _ ->
          (Failure, Cmd.none)


-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  div []
    [ h2 [] [ text "Guess it !" ]
    , viewInput model
    ]


viewInput : Model -> Html Msg
viewInput model =
  div []
    [ input [ placeholder "Text to reverse", value model.content, onInput Change ] []
    , div [] [ text (String.reverse model.content) ]
    ]



-- HTTP


getWord : Cmd Msg
getWord =
  Http.get
    { url = "dico.txt"
    , expect = Http.expectString GotText randomWord
    }


randomWord :
