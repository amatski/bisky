FROM public.ecr.aws/lambda/go:latest

RUN yum install -y go clang gcc g++ build-essential nodejs

# Copy function code
COPY judge_lambda .
ADD data/ data/

# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "judge_lambda" ]
