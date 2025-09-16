import axios from "axios";

const axiosInstace = axios.create({
  headers: { "Content-Type": "application/json" },
  baseURL: "https://api-nusakita.vercel.app/v2/",
});

export default axiosInstace;
