import {useState , useEffect} from 'react'

const useForm = (callback, validate) => {
    const [values, setValues] = useState({
        phone_number : '',
        password : ''
    })

    const [errors, setErrors] = useState({})
    const [isSubmitting, setIsSubmitting] = useState(false);

    const handleChange = e => {
        const {name, value} = e.target
        setValues({
            ...values,
            [name] : value
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
                const requestOptions = {
                    method : 'POST',
                    headers : {'Content-Type': 'application/json'},
                    body : JSON.stringify({
                        phone_number: values.phone_number,
                        password: values.password
                    })
                };
                console.log(requestOptions);            
                callback();
            }
        },
        [errors]
    )

    return {handleChange, values, handleSubmit, errors}
};

export default useForm;