// Get input + focus
let durationElement = document.getElementById("duration");
durationElement.focus();

// Runs for the duration of that app without blocking the GUI
// for live updates.
function nonBlockingIncrement(){
  function loop () {
    window.go.main.App.Clicks().then(function(clicks) {
      document.getElementById("result").innerText = `${clicks} clicks`;
    });
    window.go.main.App.ClicksPerSecond().then(function(display) {
      document.getElementById("settings").innerText = display;
    });
    (window.requestAnimationFrame || window.setTimeout)(loop);
  }
  
  loop();
}

// Setup the stop function
window.listener = function() {
  nonBlockingIncrement();
};

// Setup
window.setDelay = function() {
  window.go.main.App.SetDelay(durationElement.value);
};
