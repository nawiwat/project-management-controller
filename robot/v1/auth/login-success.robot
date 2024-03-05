*** Settings ***
Library         RequestsLibrary
Library         SeleniumLibrary
Variables       ../../config/config.yaml
Force Tags      regression    integration_test    smoke_test

*** Test Cases ***
TC_PL_0001 Login Success
    ${body}    Create Dictionary
        ...    username=Robot
        ...    password=PASSWORD1111
    ${response}=    POST   ${base_url}/login
    ...    json=${body} 
    ...    expected_status=any

    should be equal as strings    ${response.status_code}    200