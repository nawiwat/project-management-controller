*** Settings ***
Library         RequestsLibrary
Library         SeleniumLibrary
Variables       ../../config/config.yaml
Force Tags      regression    integration_test    smoke_test

*** Test Cases ***
TC_PL_0002 Login Failed By Invalid Username
    ${body}    Create Dictionary
        ...    username=invalid_username
        ...    password=PASSWORD1111
    ${response}=    POST   ${base_url}/login
    ...    json=${body} 
    ...    expected_status=any

    should be equal as strings    ${response.status_code}    400
    should be equal as strings    ${response.json()["code"]}    BAD_REQUEST
    should be equal as strings    ${response.json()["reason"]}    invalid Username

TC_PL_0003 Login Failed By Invalid Password
    ${body}    Create Dictionary
        ...    username=Robot
        ...    password=invalid_password
    ${response}=    POST   ${base_url}/login
    ...    json=${body} 
    ...    expected_status=any

    should be equal as strings    ${response.status_code}    400
    should be equal as strings    ${response.json()["code"]}    BAD_REQUEST
    should be equal as strings    ${response.json()["reason"]}    invalid password