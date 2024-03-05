*** Settings ***
Library         RequestsLibrary
Library         SeleniumLibrary
Variables       ../../config/config.yaml
Force Tags      regression    integration_test    smoke_test

*** Test Cases ***
TC_PL_0004 Create User Failer By Duplicate Username
    ${body}    Create Dictionary
        ...    username=Robot
        ...    password=PASSWORD1111
        ...    name=robot
        ...    surname=automate
        ...    email=Test@Robot.com
        ...    github=https://github.com/Robot
        ...    phone=0100100101
        ...    description=robot test account
        
    ${response}=    POST   ${base_url}/user
    ...    json=${body} 
    ...    expected_status=any

    should be equal as strings    ${response.status_code}    400
    should be equal as strings    ${response.json()["code"]}    BAD_REQUEST
    should be equal as strings    ${response.json()["reason"]}    username duplicate