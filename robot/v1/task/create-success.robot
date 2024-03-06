*** Settings ***
Library         RequestsLibrary
Library         SeleniumLibrary
Variables       ../../config/config.yaml
Force Tags      regression    integration_test    smoke_test

*** Test Cases ***

TC_PL_0005 Create Task Success

    ${body}    Create Dictionary
        ...    username=TaRobot
        ...    password=1234
    ${response}=    POST   ${base_url}/login
    ...    json=${body} 
    ...    expected_status=any
    
    ${kanban}    Create Dictionary
        ...    position=${2}
        ...    column=To Do
     ${auth_headers}    Create Dictionary
        ...    Authorization
        ...    Bearer ${response.json()["token"]}
    ${body}    Create Dictionary
        ...    project_id=${9}
        ...    name=Task-Create-Test
        ...    description=Created-From-Automate-Test
        ...    start_date=0
        ...    end_date=0
        ...    kanban=${kanban}
    
        
    ${response}=    POST   ${base_url}/task
        ...    json=${body} 
        ...    headers=${auth_headers}
        ...    expected_status=any

    should be equal as strings    ${response.status_code}    200