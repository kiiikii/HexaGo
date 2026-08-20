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

  //* --- Fetching history from message endpoint
  function loadHistory() {
    $http.get('http://localhost:8080/api/v1/secure/messages?limit=50', {
      //* Passing JWT
      headers: {'Authorization': 'Bearer ' + $scope.token}
    }).then(function(response) {
      //* the API Return {data: [...]}. We reverse it so oldest on the top
      if (response.data.data) {
        $scope.mesagges = response.data.data.reverse();
      }
    }, function (error) {
      console.error("Failed to load history:", error);
    });
  }

  //* --- Delete message
  $scope.deleteMessage = function(messageId, index) {
    console.log("🛠️ Attempting to delete message with ID:", messageId);

    if (!messageId) {
      console.error("❌ Cannot delete! Message ID is empty. Refresh the page first.");
      return
    };

    $http.delete('http://localhost:8080/api/v1/secure/messages/' + messageId, {
      headers: {'Authorization': 'Bearer ' + $scope.token}
    }).then(function(response) {
      console.log("✅ Message deleted in database!");
      //! Remove it from UI
      $scope.messages.splice(index, 1);
    }, function(error) {
      console.error("❌ Delete failed:", error);
      alert("Failed to delete. Is this your message?");
    });
  };

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

      loadHistory();

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