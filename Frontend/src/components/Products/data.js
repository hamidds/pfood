import {
    ProductButton,
    ProductCard,
    ProductImg,
    ProductInfo,
    ProductPrice,
    ProductRegion,
    ProductsContainer,
    ProductsHeading,
    ProductTitle,
    ProductWrapper,
    SearchButton,
    SearchWrapper
} from "./ProductsElements";
import product1 from "../../images/product-1.jpg";

const ShoppingList = ({data, price, click, deleteHandler}) => (
    <ProductsContainer>
        <ProductsHeading>Shopping List</ProductsHeading>
        <ProductsHeading>{price}</ProductsHeading>
        <SearchWrapper>
            <SearchButton onClick={click}>{"Back"}</SearchButton>
            <SearchButton onClick={click}>{"Pay"}</SearchButton>
        </SearchWrapper>

        <ProductWrapper>
            {data.map((item, index) => {
                return (
                    <ProductCard key={index}>
                        <ProductImg src={product1} alt={item.food.name}/>
                        <ProductInfo>
                            <ProductTitle>{item.food.name}</ProductTitle>
                            {/*<ProductRestaurant>{item.restaurant.name} restaurant</ProductRestaurant>*/}
                            <ProductRegion>{item.count}</ProductRegion>
                            <ProductPrice>{item.count * item.food.price}</ProductPrice>
                            <ProductButton onClick={() => deleteHandler(item)}>Delete from cart</ProductButton>
                        </ProductInfo>
                    </ProductCard>
                );
            })}
            {data.length === 0 && <span>shopping list is empty!</span>}
        </ProductWrapper>
    </ProductsContainer>
)

export default ShoppingList

