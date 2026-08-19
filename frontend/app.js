const app = angular.module('hexagoApp', []);

app.controller('MainController', function($scope, $http) {
  $scope.isLoggedIn = false;
  $scope.email = "";
  $scope.password = "";
  $scope.token = "";
  $scope.errorMessage = "";

  $scope.messages = [];
  $scope.messageContent = "";
  let ws = null;

  function connectWebSocket() {
    //! Open the connection using the token 
    ws = new WebSocket(`ws://localhost:8080/api/v1/ws?token=${$scope.token}`);
    ws.onopen = function() {
      console.log("WebSocket Connected");
    };

    ws.onmessage = function(event) {
      const parseData = JSON.parse(event.data);

      //! $apply ensure angular updates UI
      $scope.$apply(function() {
        $scope.messages.push(parseData);
      });
    };

    ws.onerror = function(error) {
      console.error("WebSocket Error:", error);
    };

    ws.onclose = function(event) {
        console.warn(`⚠️ WebSocket closed! Code: ${event.code}, Reason: "${event.reason}", Clean: ${event.wasClean}`);
    };
  }

  $scope.login = function () {
    const payload = {
      email: $scope.email,
      password: $scope.password
    };

    $http.post('http://localhost:8080/api/v1/login', payload).then(function(response) {
      //* Success Grab token
      $scope.token = response.data.token;
      $scope.isLoggedIn = true;
      $scope.errorMessage = "";
      console.log("Logged In Successfully", $scope.token);

      connectWebSocket();

    }, function(error) {
      $scope.errorMessage = "Login failed. Check you credentials";
      console.log("Login error:", error)
    });
  };

  $scope.sendMessage = function() {
    if ($scope.messageContent.trim() !== "" && ws) {
      const messagePayload = {
        content: $scope.messageContent
      };
      ws.send(JSON.stringify(messagePayload));
      $scope.messageContent = "";
    } else {
        console.warn("⚠️ WebSocket is not open yet. Current state:", ws ? ws.readyState : "No socket");
    }
  }
});