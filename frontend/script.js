const plantCards = document.getElementById("plantCards")

function addPlant() {

    const form = document.querySelector('form')
    
    form.addEventListener("submit", (e) => {
        e.preventDefault();

        let name = document.getElementById("name").value
        let comment = document.getElementById("comment").value
        console.log("vous avez ajouté une plante");

        fetch('http://localhost:8080/plants', {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: name,
                comment: comment,
            })
        })
    //.then(res => res.json())  --> on parse la réponse de l'api
    // aujourd'hui l'api ne renvoi rien donc ca ne sert a rien de parser
    // demain, l'api retournera la plante qui sera créée
    // a ce moment, on pourra le parser et l'ajouter à notre liste de plantes
    // tant que ce n'est pas fait, il faut actualiser la page pour actualiser la liste des plantes affichées
    .then(data => console.log('Created:', data))
    .catch(err => console.error('Error:', err));
    e.target.reset();
    })
}
addPlant() 

function deletePlant(id){
    fetch(`http://localhost:8080/plants/${id}`, {
        method: 'DELETE',
    })
     .then(response => {
    if (response.ok) {
      console.log('Resource deleted successfully.');
    } else {
      console.error('Failed to delete. Status:', response.status);
    }
  })
  .catch(error => {
    console.error('Network error:', error);
  });
}

function generatePlantCards(plants) {  //rajouter la BD pour pouvoir aller chercher les plantes dedans
    plants.forEach((plant, index) => {
        console.log("plant :", plant);
        //création d'une balise pour ajouter la plante
        const plantElement = document.createElement("div");
        plantElement.classList.add("plantElement")
        
        // ajout du nom de la plante
        const name = plant.name;
        const plantName = document.createElement("div");
        plantName.classList.add("plantName");
        plantName.innerText = name;

        // ajout de la date d'ajout de la plante
        const addDateStr = plant.created_at;
        const addDate = new Date(addDateStr);  // cree un objet Date
        const formatted = addDate.toLocaleDateString("fr-FR"); // "25/10/2025"
        const plantAddDate = document.createElement("div");
        plantAddDate.classList.add("addDate");
        plantAddDate.innerText = "Ajouté le : " + formatted;

        // ajout du commentaire de la plante
        const comment = plant.comment;
        const plantComment = document.createElement("div");
        plantComment.classList.add("plantComment");
        plantComment.innerText = comment;
        
        // ajout d'un bouton pour delete la plante
        const buttonDelete = document.createElement("button")
        buttonDelete.textContent = "Supprimer";
        buttonDelete.addEventListener("click", (e) => {
            console.log("vous avez cliqué sur le bouton de la plante :" + plant.id);
            deletePlant(plant.id);
        })
        
        // On ajoute les éléments (nom, infos, photos...) à la balise plantElement
        plantElement.appendChild(plantName);
        plantElement.appendChild(plantAddDate);
        plantElement.appendChild(plantComment);
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

