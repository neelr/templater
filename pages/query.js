import { useState, useEffect } from "react"
import { Flex, Text, Heading, Image } from "rebass"
import { useRouter } from 'next/router'
import Link from "next/link"
import fetch from "isomorphic-unfetch"
import Head from "next/head"

const UserPage = props => {
    const router = useRouter()
    let [search, setSearch] = useState()
    useEffect(() => {
        fetch(`https://templater-api.hacker22.repl.co/api/query`, {
            method: "POST",
            body: router.query.q
        })
            .then(d => d.json())
            .then(d => setSearch(d))
            .catch(d => console.log(404))
    }, router.query.q)
    if (search) {
        return (
            <Flex w="100vw" flexDirection="column">
                <Head>
                    <title>Search Results for "{router.query.q}"</title>
                </Head>
                <Heading mx="auto">Search Results for "{router.query.q}"</Heading>
                <Flex mx="auto" width={["90vw", "70vw", "60vw"]} flexWrap="wrap">
                    {
                        search.map(v => (
                            <Link href={`/${v.user}/${v.template}`}>
                                <Text m="10px" p="10px" sx={{
                                    boxShadow: "lg", borderRadius: "10px", ":hover": {
                                        color: "primary",
                                        cursor: "pointer"
                                    }
                                }}>@{v.user}/{v.template}</Text>
                            </Link>
                        ))
                    }
                </Flex>
            </Flex>
        )
    } else {
        return (
            <Text>404 not found</Text>
        )
    }
}

export default UserPage;