[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/UxpU_KWG)

# ServerlessTemplate# Leetcode Questions Serverless API 
This is a serverless function deployed on GCP Cloud Function.
The purpose of this function is to fetch Leetcode questions together with the questions content. 

## How to fetch the questions?
1. Open your browser
2. Enter the following URL 
    - `OFFSET` is an integer used to skip the first `OFFSET` of rows
    - `PAGE SIZE` is an integer used to get row `OFFSET` to row `OFFSET + PAGE SIZE`
```
https://asia-southeast1-peer-preps-assignment6.cloudfunctions.net/ay2324s1-assessment6-g13-dev-getProblems?offset={OFFSET}&page-size={PAGE SIZE}
``` 

## How to develop locally?
1. Clone this repo to your local device
2. Run `yarn install` or `npm install` to install all the necessary dependencies

### How to test development locally?
1. Need to use a Unix terminal like `bash`
2. Run the following command
```
FUNCTION_TARGET=GetProblems LOCAL_ONLY=true go run cmd/main.go
```
3. Open your browser and go to `localhost:8080?offset={OFFSET}&page-size={PAGE SIZE}