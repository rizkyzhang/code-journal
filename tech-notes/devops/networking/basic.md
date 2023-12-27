## Types of computer network

- Personal Area Network (PAN)  
  Radius is usually up to 10 meter, example is wireless speaker
- Local Area Network (LAN)
- Metropolitan Area Network (MAN)
- Wide Area Network (WAN)

## Network Latency vs Throughput vs Bandwidth

![Analogy](https://www.dnsstuff.com/wp-content/uploads/2020/12/image-11.png)

## OSI vs TCP/IP

### Diagram

![Model](https://instrumentationtools.com/wp-content/uploads/2016/06/instrumentationtools.com_difference-between-tcp-ip-model-and-osi-model.png?ezimgfmt=ngcb2/notWebP)
![Differences](https://instrumentationtools.com/wp-content/uploads/2016/06/instrumentationtools.com_comparison-of-tcp-ip-and-osi-model.png?ezimgfmt=ng:webp/ngcb2)
![OSI Model](https://www.imperva.com/learn/wp-content/uploads/sites/13/2020/02/OSI-7-layers.jpg)
![TCP/IP Model](https://linuxhint.com/wp-content/uploads/2021/10/TCPIP-layers-and-their-functions-1.jpg)

### OSI Layers

1. Application Layer  
   The application layer is used by end-user software such as web browsers and email clients. It provides protocols that allow software to send and receive information and present meaningful data to users. A few examples of application layer protocols are the Hypertext Transfer Protocol (HTTP), File Transfer Protocol (FTP), Post Office Protocol (POP), Simple Mail Transfer Protocol (SMTP), and Domain Name System (DNS).
2. Presentation Layer  
   The presentation layer prepares data for the application layer. It defines how two devices should encode, encrypt, and compress data so it is received correctly on the other end. The presentation layer takes any data transmitted by the application layer and prepares it for transmission over the session layer.
3. Session Layer  
   The session layer creates communication channels, called sessions, between devices. It is responsible for opening sessions, ensuring they remain open and functional while data is being transferred, and closing them when communication ends. The session layer can also set checkpoints during a data transfer—if the session is interrupted, devices can resume data transfer from the last checkpoint.
4. Transport Layer  
   The transport layer takes data transferred in the session layer and breaks it into “segments” on the transmitting end. It is responsible for reassembling the segments on the receiving end, turning it back into data that can be used by the session layer. The transport layer carries out flow control, sending data at a rate that matches the connection speed of the receiving device, and error control, checking if data was received incorrectly and if not, requesting it again.
5. Network Layer  
   The network layer has two main functions. One is breaking up segments into network packets, and reassembling the packets on the receiving end. The other is routing packets by discovering the best path across a physical network. The network layer uses network addresses (typically Internet Protocol addresses) to route packets to a destination node.
6. Data Link Layer  
   The data link layer establishes and terminates a connection between two physically-connected nodes on a network. It breaks up packets into frames and sends them from source to destination. This layer is composed of two parts—Logical Link Control (LLC), which identifies network protocols, performs error checking and synchronizes frames, and Media Access Control (MAC) which uses MAC addresses to connect devices and define permissions to transmit and receive data.
7. Physical Layer  
   The physical layer is responsible for the physical cable or wireless connection between network nodes. It defines the connector, the electrical cable or wireless technology connecting the devices, and is responsible for transmission of the raw data, which is simply a series of 0s and 1s, while taking care of bit rate control.

## DNS

### Anatomy

- Subdomain
- Domain
- Top-Level Domain (TLD)

### Flow

1. Browser cache -> DNS cache -> Hosts file
2. Recursive DNS Server
3. Root DNS Server
4. TLD Server
5. Authoritative Name Server

### Types of DNS Records

- A record / Address record is a record that stores IPv4 Address of a domain.
- AAAA record is a record that stores IPv6 Address of a domain.
- CNAME / Alias (canonical name) record that allows us to redirect a domain to another domain.
- MX (mail exchange) record is a record that redirect email to an email server.
- NS (nameserver) record is a record that stores Authoritative Name Server location of a domain.

## Security

### Common attacks

- DNS Cache Poisoning
- Denial-of-Service (DoS) and Distributed Denial-of-Service (DDoS)
- Man-in-the-Middle

### Prevention techniques

- DNS Security Extensions (DNSSEC) is used to verify DNS record using cryptographic signatures, usually used to prevent DNS Cache Poisoning.
- Proxy server
  - Forward proxy server implemented on client side (client anonymity)
  - Reverse proxy server implemented on server side (server anonymity)
- Rate Limit to prevent DoS and DDoS
- Firewall
- HTTPS
