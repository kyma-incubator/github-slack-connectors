# Project concept for THACK 2019

Overall, this is just a concept and should be treated as such - only a suggestion
on how the team may approach the topic of creating a showcase for Kyma.

This concept consists of a few parts or stories:
- Reacting to PR/Issue comments
- Notifying on ZenHub/GitHub events in slack channels

This docs aims at describing those in a simple, familiar way - the concept is open
for a discussion and most probably will be changed during Hack Team workshops.

## Reacting to PR/Issue comments

This one is based on a simple scenario:
1. Issue / PR is commented by someone
1. GitHub API utilizes webhook to call GitHub Connector
1. GitHub Connector processes request, an event is sent to Kyma
1. Lambda function is triggered
1. Lambda calls Azure Text Analytics service
1. If the comment is marked as 'rude' or conflicting Code of Conduct a proper action should be taken (slack notification perhaps?)
1. If the comment indicates that code should be re-reviewed a proper label should be assigned to the PR
1. Based on above - Lambda calls GitHub / Slack API to resolve corresponding action

In this scenario there are two different connectors: one for GitHub and one for Slack, they both have similar set of responsibilities - each of them is communicating with corresponding API, authorizes within Application Connector and registers its services within Kyma (so that they can be further used by Lambdas).
Additionally, GitHub Connector will need to handle GitHub calls via webhooks - it will enable reacting to particular events that happen.

Functionalities that need to be covered:
- GitHub Connector:
  - Authorizing within Application Connector in Kyma
  - Registering a service
  - Utilizing GitHub webhooks to send events
- Slack Connector:
  - Authorizing within Application Connector in Kyma
  - Registering a service

## GitHub / ZenHub events notifications

This one is designed to help teams track their task flow - there are times when some of them gets 'lost' in abundance of other things on ZenHub boards.
Whole idea is based on notifying developers that a task from their backlog or any area that they are interested in moved to 'Review' column.

The concept is following:
1. Issue is moved from 'In Progress' to 'To Review' column
1. ZenHub API utilizes webhook to call ZenHub Connector
1. ZenHub Connector processes the event, decides whether an event should be sent to Kyma
1. Lambda function is triggered, Slack Connector gets called
1. Slack Connector calls Slack API to send a message to a proper channel

Like in the first scenario - two connectors are present here.

Functionalities that need to be covered:
- ZenHub Connector:
  - Authorizing within Application Connector in Kyma
  - Registering a service
  - Utilizing ZenHub webhooks to send events
- Slack Connector:
  - Authorizing within Application Connector in Kyma
  - Registering a service
