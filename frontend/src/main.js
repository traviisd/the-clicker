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
  console.log('do more stuff'); 
};

// Setup
window.setDelay = function() {
  // Get name
  let duration = durationElement.value;
  // Call App.StartClick(name)
  window.go.main.App.SetDelay(duration);
};

nameElement.onkeydown = function(e) {
  console.log(e)
  if (e.keyCode == 13) {
    window.start()
  }
}