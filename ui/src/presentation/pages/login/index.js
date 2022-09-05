import axios from "axios"
import { useState } from "react"

export function Login() {
    const [data, setData] = useState({
        email: "",
        passw: ""
    })

    const instance = axios.create({
        baseURL: 'http://localhost:5500/v1/',
        timeout: 1000,
    });

    function handleSubmit() {
        try {
            instance.post("auth/login", JSON.stringify(data))
        } catch (error) {
            console.log(error)
        }
    }

    return (
        <section className="h-screen">
            <div className="px-6 h-full text-gray-800">
                <div className="flex xl:justify-center lg:justify-between justify-center items-center flex-wrap h-full g-6">
                    <div className="xl:ml-20 xl:w-5/12 lg:w-5/12 md:w-8/12 mb-12 md:mb-0">
                        <form>
                            <input type="text" value={data.email} onChange={(e) => setData(prev => ({...prev, email: e.target.value}))} className="mb-6 form-control block w-full px-4 py-2 text-md font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-black focus:outline-none" placeholder="Email address"/>
                            <input type="password" value={data.passw} onChange={(e) => setData(prev => ({...prev, passw: e.target.value}))}  className="mb-6 form-control block w-full px-4 py-2 text-md font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-black focus:outline-none" placeholder="Password"/>
                            <div className="flex justify-between items-center mb-6">
                                <div className="form-group form-check">
                                    <input type="checkbox" className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-black checked:border-black focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer" id="rememberMe" />
                                    <label className="form-check-label inline-block text-gray-800" htmlFor="rememberMe">Remember me</label>
                                </div>
                                <a href="#!" className="text-gray-800">Forgot password?</a>
                            </div>
                            <div className="text-center lg:text-left">
                                <button onClick={() => handleSubmit()} type="button" className="w-full inline-block px-7 py-3 bg-black text-white font-medium text-sm leading-snug uppercase rounded shadow-md hover:bg-black hover:shadow-lg focus:black focus:shadow-lg focus:outline-none focus:ring-0 active:shadow-lg transition duration-150 ease-in-out">Login</button>
                                <p className="text-sm font-semibold mt-2 pt-1 mb-0">
                                    Don't have an account?
                                    <a href="#!" className="text-red-600 hover:text-red-700 focus:text-red-700 transition duration-200 ease-in-out ml-1">Register</a>
                                </p>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </section>
    )
}
