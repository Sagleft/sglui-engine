import './pace.min.js';
import './../css/loading.css';

import './../css/app.css';
import './../css/normalize.css';
import './../css/uikit.min.css';

import './uikit.min.js';
import './uikit-icons.min.js';

import '../../wailsjs/go/consts/Constants';

import {GetPages} from '../../wailsjs/go/app/application';

window.getPages = function() {
    try {
        GetPages().then((pages) => {

            if(pages.length == 0) {
                console.log("0 pages found");
                return
            }

            pages.forEach((page) => {
                console.log(page);
            });

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
