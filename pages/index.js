import React from "react"
import { Flex, Heading, Text, Button, Image } from "rebass"
import fetch from "isomorphic-unfetch"
import Link from "next/link"
import Head from "next/head"

export default class extends React.Component {
    state = {
        users: {}
    }
    componentDidMount() {
        fetch("https://templater-api.hacker22.repl.co/api/users")
            .then(res => res.json())
            .then(json => {
                console.log(json)
                this.setState({ users: json })
            })
    }
    render() {
        return (
            <Flex w="100vw" flexDirection="column" px={["20px", 0, 0, 0]}>
                <Head>
                    <title>Templater</title>
                </Head>
                <Heading mx="auto" fontSize={[4, 5, 6]}>Templater</Heading>
                <Text sx={{ textIndent: "15px" }} width={["90vw", "50vw", "40vw"]} py="10px" mx="auto">The templating platform for all your needs! It has an easy to use CLI that links to this website, so you can showcase cool templates that you made easily and without any pain! Almost no setup and removes the need to clone a repo, or for 100's of boilerplate commands.</Text>
                <Button as="a" href="https://github.com/neelr/templater" px="20px" fontSize={[2, 3, 3]} mx="auto" sx={{
                    ":hover": {
                        bg: "secondary",
                        cursor: "pointer"
                    }
                }}>Github</Button>
                <Heading mt="25px" mx="auto">Users</Heading>
                <Flex width={["80vw", "75vw"]} mx="auto">
                    {this.state.users ? Object.values(this.state.users).map((v, i) => (
                        <Link href={`/${Object.keys(this.state.users)[i]}`}>
                            <Flex sx={{
                                boxShadow: "lg",
                                minWidth: "200px",
                                height: "100px",
                                boxShadow: "0 4px 6px -1px",
                                color: "primary",
                                borderRadius: "10px",
                                ":hover": {
                                    cursor: "pointer",
                                    color: "secondary"
                                },
                                m: "10px"
                            }}>
                                <Image m="10px" sx={{ borderRadius: "500px" }} my="auto" src={v.avatar} height="75%" />
                                <Text fontWeight={700} my="auto">@{Object.keys(this.state.users)[i]}</Text>
                            </Flex>
                        </Link>
                    )) : null}
                </Flex>
            </Flex>
        )
    }

}