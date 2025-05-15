/**
 * 
 */
class GraphQLQueryBuilder {
    name = "";
    args = {};
    scheme = {};

    /**
     * 
     * @param {string} name 
     * @param {{string:string|number}} args 
     * @param {[]} scheme 
     */
    constructor(name, args, scheme) {
        this.name = name;
        this.args = args;
        this.scheme = scheme;
    }

    /**
     * 
     * @returns {string}
     */
    build() {
        let args = this.buildArgs();

        args = `${this.name}${args ? ` ${args}` : ``}`;

        let body = this.buildScheme(this.scheme, 2);

        return `{\n${this.space()}"query": "{\\n${this.space()}${args} {\\n${body}${this.space()}\\n${this.space()}}\\n}"\n}`;
    }

    /**
     * 
     * @returns {string}
     */
    buildArgs() {
        let args = Object.keys(this.args).map(
            k => typeof this.args[k] == "string" ? `${k}:\\"${this.args[k]}\\"` : `${k}:${this.args[k]}`
        );
        return args.length ? `(${args.join(",")})` : ``;
    }

    /**
     * 
     * @param {[]} scheme 
     * @param {number} index 
     * @returns {string}
     */
    buildScheme(scheme, index = 0) {
        return scheme.map(key => {
            if (typeof key == "object") {
                return Object.keys(key).map(
                    k => `${this.space(index)}${k} {\\n${this.buildScheme(key[k], index+1)}\\n${this.space(index)}}`
                ).join(",\\n")
            }
            return `${this.space(index)}${key}`;
        }).join(",\\n");
    }

    /**
     * 
     * @param {number} dept 
     * @returns {string}
     */
    space(dept = 1) {
        let space = "  ";
        let str = "";

        for(let i =0; i < dept; i++) {
            str += space;
        }

        return str;
    }

}

/**
 * 
 */
class GraphQL {
    name = "";
    url = "";
    headers = {};
    args = {};
    structure = {};

    /**
     * 
     * @param {string} name 
     */
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
        const req = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
            body: new GraphQLQueryBuilder(this.name, this.args, this.structure).build(),
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
