"use strict";

import axios from "axios";


export const HTTP = axios.create({
    baseURL: `http://jsonplaceholder.typicode.com/`,
    headers: {
        Authorization: 'Bearer {token}'
    }
})
