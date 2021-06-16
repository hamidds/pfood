import axios from 'axios';
import {useEffect, useState} from 'react'
import {Redirect} from "react-router-dom";

const useFormCustomerRegister = (callback, validate) => {
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

    const handleSubmit = e => {
        e.preventDefault();
        setErrors(validate(values));
        setIsSubmitting(true);

    }

    useEffect(
        () => {
            if (Object.keys(errors).length === 0 && isSubmitting) {
                let body = {
                    "phone_number": values.phone_number,
                    "password": values.password
                }
                const requestOptions = {
                    body: {
                        "phone_number": values.phone_number,
                        "password": values.password
                    }
                };
                console.log(values.phone_number);
                console.log(values.password);
                console.log(requestOptions);

                // const res =  fetch('https://localhost:8000/signup/user', requestOptions);
                // axios.post('http://localhost:8000/signup/user', body).then((response) => {
                //     console.log(response.data);
                //     console.log(response.status);
                //     return <Redirect to="/main"/>
                // }).catch(function (error) {
                //     // handle error
                //     console.log(error.response.data.errors.body);
                //     errors.phone_number = error.response.data.errors.body;
                //     setErrors(errors)
                // });
                callback();
            }
        },
        [errors]
    )

    return {handleChange, values, handleSubmit, errors}
};

export default useFormCustomerRegister;