const cde = require('cdecode/client');

console.log("FORCE PUBLISH! API: ", cde);
module.exports = {
    activate: () => {
        console.log("ACTIVATING! API: ", cde);
    },
    deactivate: () => {
        console.log("DEACTIVATING :(");
    },
};
