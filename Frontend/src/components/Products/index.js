import React, {useState} from 'react';
import {
    Input,
    ProductButton,
    ProductCard,
    ProductImg,
    ProductInfo,
    ProductPrice,
    ProductRegion,
    ProductRestaurant,
    ProductsContainer,
    ProductsHeading,
    ProductTitle,
    ProductWrapper,
    SearchButton,
    SearchWrapper
} from './ProductsElements';
import axios from "axios";

import product1 from '../../images/product-1.jpg';

const apiUrl = 'http://localhost:8000';

axios.interceptors.request.use(
    config => {
        const {origin} = new URL(config.url);
        const allowedOrigins = [apiUrl];
        const token = localStorage.getItem('token');
        if (allowedOrigins.includes(origin)) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);


const Products = ({heading, data}) => {
    const [searchTextName, setSearchTextName] = useState('');
    const [searchTextRestaurant, setSearchTextRestaurant] = useState('');
    const [searchTextRegion, setSearchRegion] = useState('');
    const [searchedData, setSearchedData] = useState([]);
    const storedJwt = localStorage.getItem('token');
    const [jwt, setJwt] = useState(storedJwt || null);
    const [shoppingListPrice, setShoppingListPrice] = useState(0)
    const [shoppingList, setShoppingList] = useState([]);
    // const ss = getFoods()

    function getFoods() {

        axios.get(`${apiUrl}/foodss`).then((response) => {
            let foods = response.data.foods;
            console.log(typeof foods);
            data = foods
            setSearchedData(foods)
        }).catch(function (error) {
        });

    }

    function handleAddToCart(food, count) {
        // setSearchTextName(value);
        console.log(food)

        let item = {
            "food": food,
            "count": count
        }
        let shp = shoppingList
        shp.push(item)
        setShoppingList(shp)
        setShoppingListPrice(shoppingListPrice + food.price * count)
        console.log(shoppingListPrice)
        console.log(shoppingList)
    }

    const handleChangeName = value => {
        setSearchTextName(value);
    }
    const handleChangeRestaurant = value => {
        setSearchTextRestaurant(value);
    }
    const handleChangeRegion = value => {
        setSearchRegion(value);
    }
    const handleSearchClick = () => {
        filterDataOnName(searchTextName);
        filterDataOnRestaurant(searchTextRestaurant);
        filterDataOnRegion(searchTextRegion);
    }

    const filterDataOnName = value => {
        const lowerCaseValue = value.toLowerCase().trim();
        if (!lowerCaseValue) {
            setSearchedData(data);
        } else {
            const filteredData = data.filter(item => {
                return Object.keys(item).some(key => {
                    if (key === 'name')
                        return item[key].toString().toLowerCase().includes(lowerCaseValue);
                })
            });
            setSearchedData(filteredData);
        }
    }
    const filterDataOnRestaurant = value => {
        const lowerCaseValue = value.toLowerCase().trim();
        if (!lowerCaseValue) {
            setSearchedData(data);
        } else {
            const filteredData = data.filter(item => {
                return Object.keys(item).some(key => {
                    if (key === 'restaurant')
                        return item[key].toString().toLowerCase().includes(lowerCaseValue);
                })
            });
            setSearchedData(filteredData);
        }
    }
    const filterDataOnRegion = value => {
        const lowerCaseValue = value.toLowerCase().trim();
        if (!lowerCaseValue) {
            setSearchedData(data);
        } else {
            const filteredData = data.filter(item => {
                return Object.keys(item).some(key => {
                    if (key === 'region')
                        return item[key].toString().toLowerCase().includes(lowerCaseValue);
                })
            });
            setSearchedData(filteredData);
        }
    }

    return (
        <ProductsContainer>
            <ProductsHeading>{heading}</ProductsHeading>
            <SearchWrapper>
                <Input
                    name="food"
                    placeholder="Enter food ..."
                    value={searchTextName}
                    onChange={e => handleChangeName(e.target.value)}
                />
                <Input
                    name="restaurant"
                    placeholder="Enter restaurant ..."
                    value={searchTextRestaurant}
                    onChange={e => handleChangeRestaurant(e.target.value)}
                />
                <Input
                    name="region"
                    placeholder="Enter region ..."
                    value={searchTextRegion}
                    onChange={e => handleChangeRegion(e.target.value)}
                />
                <SearchButton
                    onClick={handleSearchClick}
                >{"Search"}</SearchButton>

                <SearchButton
                    onClick={getFoods}
                >{"Load"}</SearchButton>
            </SearchWrapper>

            <ProductWrapper>
                {searchedData.map((product, index) => {
                    return (
                        <ProductCard key={index}>
                            <ProductImg src={product1} alt={product.name}/>
                            <ProductInfo>
                                <ProductTitle>{product.name}</ProductTitle>
                                <ProductRestaurant>{product.restaurant.name} restaurant</ProductRestaurant>
                                <ProductRegion>{product.restaurant.district}</ProductRegion>
                                <ProductPrice>{product.price}</ProductPrice>
                                <ProductButton onClick={() => handleAddToCart(product, 1)}>Add to cart</ProductButton>
                            </ProductInfo>
                        </ProductCard>
                    );
                })}
                {searchedData.length === 0 && <span>No records found to display!</span>}
            </ProductWrapper>
        </ProductsContainer>
    );
};

export default Products;
