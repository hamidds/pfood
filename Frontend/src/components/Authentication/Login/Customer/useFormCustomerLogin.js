import axios from 'axios';
import {useEffect, useState} from 'react'


const api_url = `http://localhost:8000`

const useFormCustomerLogin = (callback, validate) => {
    const [values, setValues] = useState({
        phone_number: '',
        password: ''
    })

    const [errors, setErrors] = useState({})
    const [isSubmitting, setIsSubmitting] = useState(false);

    const handleChange = e => {
        const {name, value} = e.target
        setValues({
            ...values,
            [name]: value
        })
    }

    const handleSubmit = async e => {
        e.preventDefault();
        setErrors(validate(values));
        let body = {
            "phone_number": values.phone_number,
            "password": values.password
        }
        await axios.post(`${api_url}/login/user`, body)
            .then((response) => {
                localStorage.setItem("token", response.data.user.token)
                console.log(response.data.user.token)
                console.log("User Logged in")
            }).catch(function (error) {
                // handle error
                console.log("Login Error")
                let errors = {}
                let err = error.response.data.errors.body;
                if (err.includes("password")) {
                    errors.password = error.response.data.errors.body;
                } else {
                    errors.email = error.response.data.errors.body;
                }
                setErrors(errors)
            });
        setIsSubmitting(true);

    }

    useEffect(
        () => {
            if (Object.keys(errors).length === 0 && isSubmitting) {
                callback();
            }
        },
        [errors]
    )

    return {handleChange, values, handleSubmit, errors}
};

export default useFormCustomerLogin;