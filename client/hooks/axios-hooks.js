import useAxios, { configure } from 'axios-hooks'
import axios from 'axios'

const instance = axios.create({
  withCredentials: true,
  baseURL: API_SERVER,
})

configure({ instance })
