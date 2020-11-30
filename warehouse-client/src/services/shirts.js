import axios from 'axios'

const baseURL = "/products/shirts"

const getAll = async () => {
    console.log(baseURL)
    const res = await axios.get(baseURL)
    return res.data
}

const shirtsService = {
    getAll
}

export default shirtsService
