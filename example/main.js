const cde = require('cdecode/client');

console.log("FORCE PUBLISH! API: ", cde);
module.exports = {
    activate: () => {
        const query = cde.gql`{ dummy }`;
        console.log('QUERY: ', query);

        cde.apollo.query({ query })
            .then(data => console.log('APOLLO: ', data))
            .catch(data => console.log('APOLLO.ERROR: ', data));

        console.log("ACTIVATING! API: ", cde);
    },
    deactivate: () => {
        console.log("DEACTIVATING :(");
    },
};
