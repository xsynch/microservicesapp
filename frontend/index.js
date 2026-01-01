// let menuItems = {}


// const idElement = document.createElement("id")
let menuTable = document.querySelector(".menu-content")

async function getMenuItems() {
    try {
        const response = await fetch("http://localhost:9090")
        if(!response.ok){
            throw new Error(`Response Status: ${response.status}`)
        }
        const menuHtmlBody = document.createElement("tbody")
        let menuItems = await response.json()
        for (item of menuItems){
            const menuHtmlRow = document.createElement("tr")
            const idHtml = document.createElement("td")
            const nameHtml = document.createElement("td")
            const descHtml = document.createElement("td")
            const priceHtml = document.createElement("td")
            const skuHtml = document.createElement("td")
            idHtml.textContent = item.id 
            nameHtml.textContent = item.name
            descHtml.textContent = item.description 
            priceHtml.textContent = item.price
            skuHtml.textContent = item.sku
            menuHtmlRow.appendChild(idHtml)
            menuHtmlRow.appendChild(nameHtml)
            menuHtmlRow.appendChild(descHtml)
            menuHtmlRow.appendChild(priceHtml)
            menuHtmlRow.appendChild(skuHtml)
            console.log(`ID: ${item.id}`)
            console.log(`Name: ${item.name}`)
            console.log(`Description: ${item.description}`)
            menuHtmlBody.appendChild(menuHtmlRow)
        }
        menuTable.appendChild(menuHtmlBody)
        
    } catch (error){
        console.error(error)
    }

    
}

getMenuItems()

