#!/usr/bin/env python3
"""
A simple worker script for featherweight.
"""
import socket
import sys

if len(sys.argv) > 1:
  port: int = int(sys.argv[1])
else:
  port = 5000

HOST: str = "localhost"

def handle_client(sock: socket.socket) -> None:
  """
  Handling a connection with the client (the worker is a server).
  E.G. keeping the connection open, listening for requests and handling them.
  """
  while True:
    data: bytes = sock.recv(1024)
    if not data:
      break
    
    # Processing the request.
    # In this case, converting the message string to uppercase.
    response: str = data.decode('utf-8').upper()
    # After decoding the input message and converting to upper case we sent it 
    # as a response after encoding it again.
    sock.sendall(response.encode('utf-8'))

def main() -> None:
  """
  Entry point for the worker module.
  """


if __name__ == "__main__":
  main()