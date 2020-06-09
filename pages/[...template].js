/** @jsx jsx */
import { jsx } from 'theme-ui'
import { useState, useEffect } from "react"
import { Flex, Text, Heading, Image } from "rebass"
import { useRouter } from 'next/router'
import fetch from "isomorphic-unfetch"
import marked from "marked"
import dompurify from "dompurify"
import Head from "next/head"
import Link from "next/link"

const Templates = props => {
    const router = useRouter()
    const [templateData, setData] = useState();
    useEffect(() => {
        if (router.query.template) {
            fetch(`https://templater-api.hacker22.repl.co/api/templates/${router.query.template[0]}/${router.query.template[1]}`)
                .then(d => d.json())
                .then(d => setData(d))
                .catch(d => console.log(404))
        }
    }, router.query.template)
    if (templateData) {
        return (
            <Flex w="100vw" flexDirection="column">
                <Head>
                    <title>@{`${router.query.template[0]}/${router.query.template[1]}`}</title>
                </Head>
                <Heading mx="auto">
                    <Link href={`/${router.query.template[0]}`}>
                        <span sx={{
                            ":hover": {
                                color: "secondary",
                                cursor: "pointer"
                            }
                        }}>@{router.query.template[0]}</span></Link>/{`${router.query.template[1]}`}</Heading>
                <Text my="10px" mx="auto"><a sx={{
                    color: "primary",
                    ":hover": {
                        color: "secondary",
                        cursor: "pointer",
                        textDecorationStyle: "wavy"
                    }
                }} href={`https://templater-api.hacker22.repl.co/api/templates/${router.query.template[0]}/${router.query.template[1]}/download`}>Download this template</a> or use the command <code>plate get {`${router.query.template[0]}/${router.query.template[1]}`}</code></Text>
                <Heading fontSize={3} mx="auto" my="10px">Files</Heading>
                <Flex sx={{
                    p: "10px",
                    mt: "10px",
                    width: ["90vw", "75vw", "60vw"],
                    mx: "auto",
                    bg: "muted",
                    flexDirection: "column",
                }}
                    dangerouslySetInnerHTML={{ __html: templateData.files ? templateData.files.join("<br/>") : "" }} />
                <Flex sx={{
                    width: ["90vw", "75vw", "60vw"],
                    mx: "auto",
                    "a": {
                        color: "primary",
                        ":hover": {
                            color: "secondary",
                            textDecorationStyle: "wavy"
                        }
                    }
                }} flexDirection="column" dangerouslySetInnerHTML={{ __html: dompurify.sanitize(marked(templateData.README)) }} />
            </Flex>
        )
    } else {
        return (
            <Text>404 not found</Text>
        )
    }
}

export default Templates;