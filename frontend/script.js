



const plantCards = document.getElementById("plantCards")
const Corps = document.querySelector(".corps")
//console.log(plantCards)

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

function generatePlantCards(plants) {  //rajouter la BD pour pouvoir aller chercher les plantes dedans
    plants.forEach((plant, index) => {
        //création d'une balise pour ajouter la plante
        const plantElement = document.createElement("div");
        plantElement.classList.add("plantElement")
        
        // On crée les différents éléments et on ajoute les informations
        let name = plant.name;
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
    });
}  

async function loadPlants() {
  try {
    const response = await fetch('http://localhost:8080/plants');
    if (!response.ok) {
      throw new Error('Network response was not ok ' + response.statusText);
    }

    const plants = await response.json();
    console.log('Fetched data:', plants);
    return plants;
  } catch (error) {
    console.error('Fetch error:', error);
  }
}

(async () => {
    const plants = await loadPlants();
    console.log('Plants after fetch:', plants);
    console.log(plants);
    generatePlantCards(plants);
    console.log(plantCards);
})();

