import axios from 'axios'

const baseURL = "http://localhost:5000/products/shirts"

const getAll = async () => {
    console.log(baseURL)
    const res = await axios.get(baseURL)
    return res.data
}

const shirtsService = {
    getAll
}

export default shirtsService
