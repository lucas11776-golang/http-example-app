
class GraphQLQueryBuilder {
    args = {};
    scheme = {};

    /**
     * 
     * @param {*} args 
     * @param {*} scheme 
     */
    constructor(args, scheme) {
        this.args = args;
        this.scheme = scheme;
    }

    /**
     * 
     * @returns {string}
     */
    build() {
        return "";
    }


    generate(index = 0) {

    }
}


class GraphQL {
    name = "";
    url = "";
    headers = {};
    args = {};
    structure = {};

    constructor(name) {
        this.name = name;
    }

    /**
     * 
     * @param {{string:string|number}} args 
     * @return {GraphQL}
     */
    arguments(args) {
        this.args = args;
        return this;
    }

    /**
     * 
     * @param {[]} scheme
     * @return {GraphQL}
     */
    scheme(scheme) {
        this.structure = scheme;
        return this;
    }

    /**
     * 
     * @param {string} url 
     * @return {Promise}
     */
    async fetch(url) {
        const filtersKeys = Object.keys(this.args);
        const filtersString = filtersKeys.map(k => {
            if (typeof this.args[k] == "string") {
                return `${k}: \"${this.args[k]}\"`;
            }
            return `${k}:${this.args[k]}`;
        });



        let s = new String(
            `{${this.name} ${filtersKeys.length ? `(${filtersString.join(",")})` : ``} {` +
            `${this.structure.map(key => `    ${key}`).join(",")}` +
            `}`
        ).toString()



        let query = `{\n  "query": "${s}  }"\n}`;

        console.log(query)


        // const req = await fetch(url, {
        //     method: 'POST',
        //     headers: {
        //         'Content-Type': 'application/json',
        //         'Accept': 'application/json',
        //     },
        //     body: query,
        // });




        // // {"query":"{\n    articles {\n        title\n    }\n}","variables":{}}


        // return await req.json();
    }
}


/**
 * Testing GraphQL basic query builder.
 * 
 * @param {string} name
 * @param {{string:string|number}}  
 * @param {[]} scheme
 * @returns {GraphQL}
 */
const Request = (name) => new GraphQL(name);


const req = Request("articles")
        // .arguments({category: "business"})
        .scheme(['title'])
        .fetch(`http://127.0.0.1:8080/graph_ql/news`)
        .then()
