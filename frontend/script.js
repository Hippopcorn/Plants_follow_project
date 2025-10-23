const plantCards = document.getElementById("plantCards")
const Corps = document.querySelector(".corps")
console.log(plantCards)

function addPlant() {
    const buttonAdd = document.createElement("button");
    buttonAdd.classList.add("buttonAdd");
    buttonAdd.textContent = "Ajouter une nouvelle plante";
    buttonAdd.addEventListener("click", (e) => {
        console.log("vous avez cliqué sur le bouton");
    })
    Corps.appendChild(buttonAdd)
}
addPlant() 

function generatePlantCards(lengthBD) {  //rajouter la BD pour pouvoir aller chercher les plantes dedans
    let nbrPlants = lengthBD
    for (let i= 0; i < nbrPlants; i++) {  // nbrPlants a definir en fonction de la taille de la BD
           
            //création d'une balise pour ajouter la plante
            const plantElement = document.createElement("div");
            plantElement.classList.add("plantElement")
            
            // On crée les différents éléments et on ajoute les informations
            let name = "Monstera"
            const plantName = document.createElement("div");
            plantName.classList.add("plantName");
            plantName.innerText = name;
      
            let addDate = "12/12/2024"
            const plantAddDate = document.createElement("div");
            plantAddDate.classList.add("addDate");
            plantAddDate.innerText = addDate;
      
            const buttonDelete = document.createElement("button")
            buttonDelete.textContent = "Supprimer";
            buttonDelete.addEventListener("click", (e) => {
                console.log("vous avez cliqué sur le bouton");
            })
            

            // On ajoute les éléments (nom, infos, photos...) à la balise plantElement
            plantElement.appendChild(plantName);
            plantElement.appendChild(plantAddDate);
            plantElement.appendChild(buttonDelete);
      
            // On ajoute plantElement dans la div plantCards
            plantCards.appendChild(plantElement);          
    }
}   
generatePlantCards(50)
console.log(plantCards)