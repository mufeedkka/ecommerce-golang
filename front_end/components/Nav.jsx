"use client";

import Link from "next/link";
import Image from "next/image";
import { useState, useEffect } from "react";

const Nav = () => {
  return (
    <nav className="flex-between w-full mb-16 pt-3">
        <Link href="/" className="flex gap-2 flex-center">
            <Image src="/assets/images/logo.svg" alt="logo" width={30} height={30} className="object-contain"/>
        <p className="logo_text">eGramShop</p>
        </Link>
        <Link href="/" className="flex gap-2 flex-right">
          Home
        </Link>
        <Link href="/about" className="flex gap-2 flex-right">
          About
        </Link>
        <Link href="/products" className="flex gap-2 flex-right">
          Products
        </Link>
    </nav>
  )
}

export default Nav