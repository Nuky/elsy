Feature: templates

  Scenario: with an invalid template
    When I run `lc --template foo bootstrap`
    Then it should fail
    And the output should contain 'template \"foo\" does not exist'
