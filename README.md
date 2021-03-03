# Gopher Translation Service

## Developer Notes

- [x] Setup a git repo
- [x] Prepare the project directory structure
- [x] Create an empty, scaffolding project
- [x] Setup the http server gorilla mux

## Word translator

- [x] Create an endpoint accepting POST requests “/word”

- [x] Parse the JSON payload into a struct

  ```json
  {"english-word": "<a single English word>"}
  ```

- [x] Validate the input data

    - [x] If there is no payload
    - [x] If there is no word in the payload
    - [x] If there are multiple words

- [x] Return a response in JSON format

  ```json
  {"gopher-word": "<translated version of the given word>"}
  ```

    - [x] Make the word translation function
        - [x] If a word starts with a vowel letter, add prefix “g” to the word (ex. apple => gapple)
        - [x] If a word starts with the consonant letters “xr”, add the prefix “ge” to the begging of the word. Such words as “xray” actually sound in the beginning with vowel sound as you pronounce them so a true gopher would say “gexray”.
        - [x] If a word starts with a consonant sound, move it to the end of the word and then add “ogo” suffix to the word. Consonant sounds can be made up of multiple consonants, a.k.a. a consonant cluster (e.g. "chair" -> "airchogo”).
            - [x] What if the word has no vowel letters, i.e. `smth` for example
        - [x] If a word starts with a consonant sound followed by "qu", move it to the end of the word, and then add "ogo" suffix to the word (e.g. "square" -> "aresquogo").
        - [x] Don’t use words like - “don’t”, “shouldn’t”, etc. Even translated they still won’t understand you so skip them in your solution.
        - [x] preserve the case (if possible)

- [x] Store the input and output in a map for history purposes

## Sentence translator

- [x] Create an endpoint accepting POST requests “/sentence”

- [x] Parse the JSON payload into a struct

  ```json
  {"english-sentence": "<sentence of English words>"}
  ```

- [x] Validate the input

    - [x] If there is a payload
    - [x] If there are sentence and how many

- [x] Create a translation function that will utilize the word translation function from above

    - [x] Tokenize the sentence
    - [x] Extract the ending character (?,.)
    - [x] Preserve the case (if possible)
    - [x] Preserve the ending character (?,.)

- [x] Store the input and output in a map for history purposes



## History

- [x] Create a new endpoint GET “/history”

- [x] Create a map holding the translated words/sentences

- [x] Sort the map by its keys alphabetically

- [x] Serialize it to jSON and return in the following format:

  ```json
  {
    "history": [
      {
        "apple": "gapple"
      },
      {
        "my": "ymogo"
      }
    ]
  }
  ```



## Testing

- [x] Write unit tests to ensure there is no bugs
    - [x] Test the request handlers
        - [x] Test with different payload
        - [x] Test without payload
        - [x] Verify the status code
        - [x] Verify the content type
        - [x] Do basic output check
    - [x] Test the translation functions
        - [x] Different combination of  letters (vowel, consonant, xr and so on)
        - [x] Different case (title case, lower case, upper case)
        - [x] Ensure that words like - “don’t”, “shouldn’t” are not translated
        - [x] Test with sentences ending with dot, question or exclamation mark
        - [x] Test with punctuations and multiple spaces in the sentences.



## Installation

Build & run

```bash
go get
go build
./gopher-translator
```

Then open your browser on http://127.0.0.1:10000

Run the test with

```bash
go test -v ./...
```

