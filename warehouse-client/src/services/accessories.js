import axios from 'axios'

const baseURL = "http://localhost:5000/products/accessories"

const getAll = async () => {
    console.log(baseURL)
    const res = await axios.get(baseURL)
    return res.data
}

const accessoriesService = {
    getAll
}

export default accessoriesService
