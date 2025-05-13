

class GraphQL {
    name;
    url;
    args;
    structure;
    headers;

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
                return `${k}: "${this.args[k]}"`;
            }
            return `${k}:${this.args[k]}`;
        });

        let query = `{\r\n  ${this.name} ${filtersKeys.length ? `(${filtersString.join(",")})` : ``} {\r\n`;
        query += `${this.structure.map(key => `    ${key}`).join(",\r\n")}`
        query += `\r\n  }\r\n}`;

        const req = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
            body: query,
        });

        return await req.json();
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
