import './../css/app.css';
import './pace.min.js';
import './../css/loading.css';

import './../css/normalize.css';
import './../css/uikit.min.css';

import './uikit.min.js';
import './uikit-icons.min.js';

import {GetPages} from '../../wailsjs/go/app/application';

// show app container after loading
document.addEventListener("DOMContentLoaded", function(event) { 
    document.getElementById('app-container').style.display = 'block';
});

function handlePages(pages) {
    if(pages == undefined) {
        console.log("pages is not set");
        return
    }

    if(pages.length == 0) {
        console.log("0 pages found");
        return
    }

    pages.forEach((page) => {
        handlePage(page);
    });
}

function handlePage(page) {
    // TODO: find main page by tag

    page.elements.forEach((element) => {
         console.log("handle element", element);
         handleElement(element);
    });
}

function handleElement(element) {
    if(element.type == "sidebar") {
        // create sidebar
        var sidebar = document.createElement('div');
        sidebar.className = 'sidebar';
        document.body.appendChild(sidebar);

        // create sidebar title
        if(element.data.title != '') {
            var sidebarTitle = document.createElement('h4');
            sidebarTitle.innerHTML = element.data.title;
            sidebar.appendChild(sidebarTitle);
        }

        // create list
        var sidebarList = document.createElement('ul');
        sidebarList.className = 'uk-list uk-list-divider';
        sidebar.appendChild(sidebarList);

        // add buttons
        element.data.buttons.forEach((button) => {
            var btn = document.createElement('button');
            btn.className = 'uk-button uk-button-primary uk-border-rounded';
            btn.innerHTML = button.data.label;
            sidebarList.appendChild(btn);
        });
    }
}

window.getPages = function() {
    try {
        GetPages().then((pages) => {
            handlePages(pages);
        }).catch((err) => {
            console.error("handle result", err);
        });
    } catch (err) {
        console.error("try request", err);
    }
}

//let nameElement = document.getElementById("name");
//nameElement.focus();
//let resultElement = document.getElementById("result");

// Setup the greet function
/*window.greet = function () {
    // Get name
    let name = nameElement.value;

    // Check if the input is empty
    if (name === "") return;

    // Call App.Greet(name)
    try {
        Greet(name)
            .then((result) => {
                // Update result with data back from App.Greet()
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
*/
