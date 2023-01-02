/** @jsx jsx */
import { jsx, useColorMode } from 'theme-ui'
import { useState, useEffect } from "react"
import { Flex, Text, Heading, Image } from "rebass"
import { useRouter } from 'next/router'
import fetch from "isomorphic-unfetch"
import marked from "marked"
import dompurify from "dompurify"
import Head from "next/head"
import Link from "next/link"
import { CssBaseline, GeistProvider, Tree } from '@geist-ui/react'

const Templates = props => {
    const router = useRouter()
    const [templateData, setData] = useState();
    const [colorMode, setColorMode] = useColorMode()
    const [notFound, setFound] = useState("");
    useEffect(() => {
        if (router.query.template) {
            fetch(`https://plate.neelr.dev/api/templates/${router.query.template[0]}/${router.query.template[1]}`)
                .then(d => d.json())
                .then(d => setData(d))
                .catch(d => setFound("404 Not Found"))
        }
    }, router.query.template)
    if (templateData) {


        let result = [];
        let level = { result };

        templateData.files.forEach(path => {
            path.split('/').reduce((r, name, i, a) => {
                if (!r[name]) {
                    r[name] = { result: [] };
                    r.result.push({ name, children: r[name].result })
                }

                return r[name];
            }, level)
        })
        //result = result[0].children
        let mapFiles = f => {
            if (!f.name) {
                console.log(f)
                if (f.children.length > 0) {
                    return f.children.map(mapFiles)
                }
                return null
            }
            if (f.name.includes(".")) {
                return <Tree.File name={f.name} />
            } else {
                return <Tree.Folder name={f.name}>{f.children.map(mapFiles)}</Tree.Folder>
            }
        }

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
                }} href={`https://plate.neelr.dev/api/templates/${router.query.template[0]}/${router.query.template[1]}/download`}>Download this template</a> or use the command <code>plate get {`${router.query.template[0]}/${router.query.template[1]}`}</code></Text>
                <Heading fontSize={3} mx="auto" my="10px">Files</Heading>
                <Flex sx={{
                    p: "10px",
                    mt: "10px",
                    width: ["90vw", "75vw", "60vw"],
                    mx: "auto",
                    bg: "muted",
                    flexDirection: "column",
                }}
                >
                    <GeistProvider theme={{ type: colorMode !== "dark" ? "light" : "dark" }}>
                        <CssBaseline />
                        <Tree>
                            {mapFiles(result[0])}
                        </Tree>
                    </GeistProvider>
                </Flex>
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
            <Heading mx="auto" my="20px" color="primary">{notFound}</Heading>
        )
    }
}

export default Templates;
