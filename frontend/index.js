let menuItems = ""
let menusection = document.querySelector(".menu-contents")
menusection.textContent = "hello"

const idElement = document.createElement("id")

async function getMenuItems() {
    menuItems = await fetch("http://localhost:9090")
    console.log(menuItems)

    
}

