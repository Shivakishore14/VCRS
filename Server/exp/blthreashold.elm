port module Main exposing (..)

import Html exposing (Html, div, button, input, h2, text, label)
import Html.Attributes exposing (class, id, value, type_)
import Html.Events exposing (onClick, onInput)
import Json.Decode exposing (list, int, string, float, nullable, Decoder)
import Json.Encode
import Json.Decode.Pipeline exposing (decode, required, optional, hardcoded)
import Http


main : Program Never Model Msg
main =
    Html.program
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }


type alias Model =
    { levels : List Input
    , cnt : Int
    , message : String
    , testName : String
    }


type alias Input =
    { number : Int
    , val : Int
    }


type Msg
    = NoOp
    | LevelVal Input String
    | GetLevelsJson (Result Http.Error (List Input))
    | Test (Result Http.Error String)
    | BtnSubmit
    | SetTestName String


init : ( Model, Cmd Msg )
init =
    ( initModel, getJson initModel )


initModel : Model
initModel =
    { levels = initInputs, cnt = 4, message = "", testName = "elmTest" }


initInputs : List Input
initInputs =
    generateInputs 4 []


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        NoOp ->
            ( model, Cmd.none )

        LevelVal ip val ->
            ( { model | levels = List.map (changeInput val ip) model.levels }, Cmd.none )

        GetLevelsJson (Ok newList) ->
            ( { model | levels = newList }, Cmd.none )

        GetLevelsJson (Err e) ->
            ( { model | message = toString e }, Cmd.none )

        BtnSubmit ->
            ( model, putJson model )

        Test _ ->
            model ! []

        SetTestName newTest ->
            let
                newModel =
                    { model | testName = newTest }
            in
                ( newModel, getJson newModel )


port setTestName : (String -> msg) -> Sub msg


subscriptions : Model -> Sub Msg
subscriptions model =
    setTestName SetTestName


changeInput : String -> Input -> Input -> Input
changeInput val newip oldip =
    let
        ip =
            if newip.number == oldip.number then
                Input newip.number (Result.withDefault oldip.val (String.toInt val))
            else
                oldip
    in
        ip
            |> validateInputVal


validateInputVal : Input -> Input
validateInputVal ip =
    let
        val =
            ip.val

        newval =
            if val >= 0 then
                (if val <= 100 then
                    val
                 else
                    100
                )
            else
                0
    in
        Input ip.number newval


getJson : Model -> Cmd Msg
getJson model =
    let
        url =
            "/level/?testName=" ++ model.testName

        request =
            Http.get url inputListDecode
    in
        Http.send GetLevelsJson request


urlDecoder : Decoder String
urlDecoder =
    Json.Decode.at [] Json.Decode.string


encodeInput : Input -> Json.Encode.Value
encodeInput input =
    Json.Encode.object
        [ ( "number", Json.Encode.int input.number )
        , ( "val", Json.Encode.int input.val )
        ]


encodeModel : Model -> Json.Encode.Value
encodeModel model =
    Json.Encode.object
        [ ( "levels", Json.Encode.list <| List.map encodeInput model.levels )
        , ( "testName", Json.Encode.string model.testName )
        ]


encodeModelString : Model -> String
encodeModelString model =
    toString model


putJson : Model -> Cmd Msg
putJson model =
    let
        url =
            "/level/"

        body =
            model
                |> encodeModel
                |> Http.jsonBody

        request =
            Http.post url body string
    in
        Http.send Test request



--view


view : Model -> Html Msg
view model =
    div []
        [ div []
            (List.map (generateHtmlInput model) model.levels)
        , div [] [ text model.message ]
        , button [ class "btn btn-success", onClick BtnSubmit ] [ text "Submit" ]
        , div [] [ text <| toString model ]

        --, input [ value model.vals, onInput Level1 ] []
        ]


generateHtmlInput : Model -> Input -> Html Msg
generateHtmlInput model ip =
    div [ class "form-group" ]
        [ label [] [ text <| "Level" ++ toString ip.number ]
        , input [ class "form-control", value <| toString ip.val, onInput (LevelVal ip) ] []
        ]


generateInputs : Int -> List Input -> List Input
generateInputs i inputList =
    let
        temp =
            if i > 0 then
                [ Input i 0 ]
            else
                []

        outList =
            inputList ++ temp
    in
        if i > 0 then
            generateInputs (i - 1) outList
        else
            outList


inputDecoder : Decoder Input
inputDecoder =
    decode Input
        |> required "number" int
        |> required "val" int


inputListDecode : Decoder (List Input)
inputListDecode =
    list inputDecoder
