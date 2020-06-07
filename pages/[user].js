import { useState, useEffect } from "react"
import { Flex, Text, Heading, Image } from "rebass"
import { useRouter } from 'next/router'
import Link from "next/link"
import fetch from "isomorphic-unfetch"
import Head from "next/head"

const UserPage = props => {
    const router = useRouter()
    const [userData, setData] = useState();
    useEffect(() => {
        fetch(`https://templater-api.hacker22.repl.co/api/user/${router.query.user}`)
            .then(d => d.json())
            .then(d => setData(d))
            .catch(d => console.log(404))
    }, router.query.user)
    if (userData) {
        return (
            <Flex w="100vw" flexDirection="column">
                <Head>
                    <title>@{router.query.user}</title>
                </Head>
                <Image sx={{
                    height: "150px",
                    width: "150px",
                    mb: "10px",
                    mx: "auto",
                    borderRadius: "150px"
                }} src={userData.avatar} />
                <Heading mx="auto">@{router.query.user}</Heading>
                <Text mx="auto" textAlign="center" width={["90vw", "50vw", "30vw"]}>{userData.bio}</Text>
                <Heading my="10px" fontSize={3} mx="auto">Templates</Heading>
                <Flex mx="auto" width={["90vw", "70vw", "60vw"]} flexWrap="wrap">
                    {
                        userData.templates ?
                            userData.templates.map(v => (
                                <Link href={`/${router.query.user}/${v}`}>
                                    <Text m="10px" p="10px" sx={{
                                        boxShadow: "lg", borderRadius: "10px", ":hover": {
                                            color: "primary",
                                            cursor: "pointer"
                                        }
                                    }}>@{router.query.user}/{v}</Text>
                                </Link>
                            )) : <Text mx="auto">No Templates!</Text>
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