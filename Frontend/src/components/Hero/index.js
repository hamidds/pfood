import React, {useState} from 'react';
// import Navbar from '../Navbar';
import Sidebar from '../Sidebar';
import Navbar from '../Navbar'
import {productData} from '../Products/data.js'
import {HeroBtn, HeroContainer, HeroContent, HeroH1, HeroItems, HeroP} from './HeroElements';
import product1 from '../../images/product-1.jpg';
import product2 from '../../images/product-2.jpg';
import product3 from '../../images/product-3.jpg';
import Products from "../Products";

const Hero = () => {
    const [isOpen, setIsOpen] = useState(false);

    const toggle = () => {
        setIsOpen(!isOpen);
    };

    const handlePlaceOrderClick = async () => {
        // const { foods } = await axios.get(`localhost:8000/foods`);
        productData = [
            {
                img: product1,
                alt: 'Pizza',
                name: 'Supreme Pizza',
                restaurant: 'Dinner Club',
                region: 'Hudson Valley',
                price: '$19.99',

                button: 'Add to Cart'
            },
            {
                img: product2,

                alt: 'Pizza',
                name: 'Hawaiian Paradise',
                restaurant: 'Kitchen Corral',
                region: 'Western New York',
                price: '$16.99',
                button: 'Add to Cart'
            },
            {
                img: product3,
                alt: 'Pizza',
                name: 'Veggie Overload',
                restaurant: 'White Napkin Delight',
                region: 'Capital District',
                price: '$14.99',
                button: 'Add to Cart'
            }
        ];
        console.log("Asdsadasdad")
        // try {
        //
        //
        // } catch (err) {
        // }
    };

    return (
        <HeroContainer>
            <Navbar toggle={toggle}/>
            <Sidebar isOpen={isOpen} toggle={toggle}/>
            <HeroContent>
                <HeroItems>
                    <HeroH1>Find your favorite food</HeroH1>
                    <HeroP>in 60 seconds</HeroP>
                    <HeroBtn onClick={handlePlaceOrderClick}>Place Order</HeroBtn>
                </HeroItems>
            </HeroContent>
        </HeroContainer>
    );
};

export default Hero;
