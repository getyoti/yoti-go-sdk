Feature: Demo

  Scenario: Should create session and get the session result
    Given I am on "/"
    And I configure the session response

    And I switch to the iframe
    And I choose "PASSPORT"
    And I click on add photo button
    And I upload a document
    And I click on finish button
    And I wait 15 seconds

    When I am on "/success"
    Then I should see "Get Session Result"

    And the authenticity check breakdown sub check should be "security_features"
    And the authenticity check breakdown result should be "NOT_AVAILABLE"

    And the text data check breakdown sub check should be "document_in_date"
    And the text data check breakdown result should be "PASS"
