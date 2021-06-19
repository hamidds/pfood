import {useEffect, useState} from 'react'
import axios from "axios";

const api_url = `http://localhost:8000`

const useFormManagerRegister = (callback, validate) => {
    const [values, setValues] = useState({
        email: '',
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
        await axios.get(`${api_url}/signup/check/email/${values.email}`)
            .then((response) => {
                // localStorage.setItem("token", response.data.user.token)
                // console.log(response.data.user.token)
                console.log("then")
            }).catch(function (error) {
                // handle error
                let errors = {}
                errors.email = error.response.data.errors.body;
                setErrors(errors)
            });
        setIsSubmitting(true);
    }

    useEffect(
        () => {
            if (Object.keys(errors).length === 0 && isSubmitting) {
                let body = {
                    "email": values.email,
                    "password": values.password
                }
                console.log(values.email);
                console.log(values.password);

                axios.post(`${api_url}/signup/manager`, body)
                    .then((response) => {
                        localStorage.setItem("token", response.data.manager.token)
                        console.log(response.data.manager.token)
                    })
                callback();
            }
        },
        [errors]
    )

    return {handleChange, values, handleSubmit, errors}
};

export default useFormManagerRegister;