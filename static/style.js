let togg1 = document.getElementById("togg1");
let d12 = document.getElementById("d12");
togg1.addEventListener("click", () => {
  if(getComputedStyle(d12).display != "flex"){
    d12.style.display = "flex";
  } else {
    d12.style.display = "none";
  }
})
function rangeSlide(value) {
  document.getElementById('rangeValue').innerHTML = value;
}
function rangeSlide2(value) {
  document.getElementById('rangeValue2').innerHTML = value;
}
function rangeSlide3(value) {
  document.getElementById('rangeValue3').innerHTML = value;
}