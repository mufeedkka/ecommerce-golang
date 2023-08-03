import React from 'react'
import Link from "next/link";


const Products = () => {
  return (
    <div>

<Link href="/product_form">
  <button className="bg-green-500 text-white py-2 px-4 rounded hover:bg-green-600">
    Create Product
  </button>
</Link>
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
    <div className="max-w-sm rounded overflow-hidden shadow-md bg-white">
      {/* <img src="product-image.jpg" alt="Product Image" class="w-full h-40 object-cover"/> */}
      <div className="p-4">
        <h3 className="text-xl font-semibold mb-2">Product Name</h3>
        <p className="text-gray-600 text-sm mb-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
        <p className="text-lg font-semibold mb-4">$19.99</p>
        <button className="bg-green-500 text-white py-2 px-4 rounded hover:bg-green-600">Add to Cart</button>
      </div>
    </div>
  </div>
    </div>
  )
}

export default Products