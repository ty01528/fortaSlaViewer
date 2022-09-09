function editAddress(address,name){
    // showElement()
    // console.log(address)
    // console.log(name)
    document.getElementById("editPageAddress").innerHTML = address
    document.getElementById("editAddress").value = address
}
function deleteSubmit(){
    // showElement()
    document.getElementById("editForm").action = "/deleteAddress"
    document.getElementById("editForm").submit()
}
function editSubmit(){
    document.getElementById("editForm").action = "/editName"
    document.getElementById("editForm").submit()
}