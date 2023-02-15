import axios from "axios";


const instance = axios.create({
    baseURL: '',
    timeout: 5000
})


export default instance