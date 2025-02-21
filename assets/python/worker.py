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

HOST: str = "127.0.0.1"

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
    sock.sendall((response + '\n').encode('utf-8'))

def main() -> None:
  """
  Entry point for the worker module.
  """
  with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as listener_socket:
    listener_socket.bind((HOST, port))
    listener_socket.listen()
    print(f"Python worker listening on {HOST}:{port}")
    while True:
      # Repeatedly listening for new connections.
      connection_socket, addr = listener_socket.accept()
      # Keeping the connection alive and handling the requests.
      handle_client(connection_socket)
      connection_socket.close()
    


if __name__ == "__main__":
  main()