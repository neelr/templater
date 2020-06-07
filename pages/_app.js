import React from "react"
import { ThemeProvider } from "theme-ui"
import { Flex } from "rebass"
import theme from "../components/theme"
import Nav from "../components/nav"
import Footer from "../components/footer"

export default ({ Component, props }) => (
    <ThemeProvider theme={theme}>
        <Flex flexDirection="column" minHeight="100vh">
            <Nav />
            <Component {...props} />
            <Footer />
        </Flex>
    </ThemeProvider>
)