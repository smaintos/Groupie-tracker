const searchinput = document.getElementById('Search_Bar');
searchinput.addEventListener('keyup', function(){
    const input = searchinput.value;
    console.log(input);
    const result = [];
    apiObj.forEach(item => {
        if (item.name.toLowerCase().includes(input.toLowerCase())) {
            console.log(item.name);
            result.push(item.name);
        } else if (item.creationDate == parseInt(input)) {
            result.push(item.creationDate);
        } else if (item.firstAlbum.toLowerCase().includes(input.toLowerCase())) {
            result.push(item.firstAlbum);
        } else {
            for (let i = 0; i < item.members.length; i++) {
                if (item.members[i].toLowerCase().includes(input.toLowerCase())) {
                    result.push(item.members[i]);
                }
            }
        }
    });
    // supprimer les doublons (dates)
    console.log(result);
    const suggestions = document.getElementById('suggestion');
    while ( suggestions.firstChild ) suggestions.removeChild( suggestions.firstChild );
    if(input != ''){
        result.forEach(resultitem => {
            let newDiv = document.createElement('div');   
            newDiv.innerHTML = resultitem;
            newDiv.classList.add('suggestion');
            suggestions.appendChild(newDiv);
        });
    }

    suggestions.addEventListener('click', function(event) {
        const target = event.target;
        searchinput.value = target.innerHTML;
        suggestions.innerHTML = '';
    });
});