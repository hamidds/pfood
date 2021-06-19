import axios from 'axios';
import {useEffect, useState} from 'react'


const api_url = `http://localhost:8000`
const useProfileForm = (callback, validate) => {
    const [values, setValues] = useState({
        phone_number: '',
        password: '',
        address: '',
        district: 0,
        name: ''
    })
    useEffect(() => {
        axios.get(`${api_url}/profile`).then((response) => {
            localStorage.setItem("token", response.data.user.token)
            console.log(response.data.user.token)
            let profile = {
                name: response.data.user.name,
                phone_number: response.data.user.phone_number,
                password: response.data.user.password,
                address: response.data.user.address,
                district: response.data.user.district
            }
            setValues(profile)

        }).catch(function (error) {
            // handle error
            console.log("profile error")
        });
        // let profile = {
        //     name: 'hamidreza',
        //     phone_number: '09215226974',
        //     password: '',
        //     address: 'Hoanrestan',
        //     district: 5
        // }
        // setValues(profile)

    }, [])


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
                    name: values.name,
                    phone_number: values.phone_number,
                    password: values.password,
                    address: values.address,
                    district: values.district
                }
                console.log(values.phone_number);
                console.log(values.password);

                axios.put(`${api_url}/profile`, body).then((response) => {
                    localStorage.setItem("token", response.data.user.token)
                    console.log(response.data.user.token)
                }).catch(function (error) {
                    // handle error
                    let errors = {}
                    errors.phone_number = error.response.data.errors.body;
                    setErrors(errors)
                });
                callback();
            }
        },
        [errors]
    )

    return {handleChange, values, handleSubmit, errors}
};

export default useProfileForm;