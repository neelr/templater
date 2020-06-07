/** @jsx jsx */
import { jsx, useColorMode } from 'theme-ui'
import React, { useState } from "react"
import { Flex, Text } from "rebass"
import { Sun, Moon, Search } from "react-feather"
import Link from "next/link"

export default () => {
    let [query, setQuery] = useState("")
    const [colorMode, setColorMode] = useColorMode()
    return (
        <Flex w="100vw" height="75px" mx="20px">
            <Link href="/">
                <Text my="auto" fontWeight={700} sx={{
                    ":hover": {
                        color: "secondary",
                        cursor: "pointer"
                    }
                }}>Home</Text>
            </Link>
            <input placeholder="Search..." sx={{
                bg: "muted",
                border: 0,
                height: "40px",
                outline: "none",
                borderRadius: "100px",
                px: "20px",
                ml: "20px",
                my: "auto",
                width: "200px",
                color: "text",
                "::placeholder": {
                    color: "text",
                    opactity: "0.6"
                }
            }} onChange={e => setQuery(e.target.value)} onKeyPress={e => {
                if (e.key == "Enter" && query.replace(" ", "") != "") {
                    window.location.href = "/query?q=" + query
                }
            }} />
            <Flex ml="auto" my="auto" width="50px">
                <Flex m="auto" p="5px" sx={{
                    transition: "all 0.2s",
                    borderRadius: "300px",
                    ":hover": {
                        color: "primary",
                        borderColor: "primary",
                        borderWidth: "2px",
                        borderStyle: "solid",
                        cursor: "pointer"
                    }
                }} onClick={() => {
                    setColorMode(colorMode === 'default' ? 'dark' : 'default')
                }}>
                    {colorMode == "default" ? <Sun size={26} /> : <Moon size={26} />}
                </Flex>
            </Flex>
        </Flex>)
}