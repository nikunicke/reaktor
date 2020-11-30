import axios from 'axios'

const baseURL = "/products/jackets"

const getAll = async () => {
    console.log(baseURL)
    const res = await axios.get(baseURL)
    return res.data
}

const jacketsService = {
    getAll
}

export default jacketsService
